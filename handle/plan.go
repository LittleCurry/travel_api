package handle

import (
	"net/http"

	"strconv"

	"git.iguiyu.com/park/api/db"
	"git.iguiyu.com/park/api/misc"
	"git.iguiyu.com/park/api/vm"
	"git.iguiyu.com/park/struct/model"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
)

func ListPlan(c echo.Context) (err error) {
	userId := c.Get("userId").(string)

	plans := make([]vm.TxPlan, 0)
	profiles := make([]model.OwnerProfile, 0)
	err1 := db.MySQL().Where("owner_user_id=?", userId).Find(&profiles)
	if err1 == nil {
		for _, profile := range profiles {
			plan := vm.TxPlan{}
			copier.Copy(&plan, &profile)
			profileDetails := make([]model.OwnerProfileDetail, 0)
			err2 := db.MySQL().Where("owner_profile_id=?", profile.Id).Find(&profileDetails)
			if err2 == nil {
				copier.Copy(&plan.ProfileDetail, profileDetails)
			}
			plans = append(plans, plan)
		}
		return c.JSON(http.StatusOK, plans)
	}
	/*
		profiledetails := make([]joinmodel.ProfileDetailProfile, 0)
		err1 := db.MySQL().Join(
			"INNER",
			"owner_profile",
			"owner_profile.id = owner_profile_detail.owner_profile_id").Where(
			"owner_profile.owner_user_id=" + userId).Find(&profiledetails)
		if err1 == nil {
			plans := vm.TxPlan{}
			copier.Copy(&plans, &profiledetails)
			return c.JSON(http.StatusOK, plans)
		}*/
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func CreatePlan(c echo.Context) (err error) {
	userId := c.Get("userId").(string)

	var plan = vm.RxPlan{}
	if err1 := c.Bind(&plan); err1 == nil {

		profile := model.OwnerProfile{}
		copier.Copy(&profile, &plan)
		profile.OwnerUserId = misc.FormatUserIdToInt(userId)

		affected, err2 := db.MySQL().Insert(&profile)
		hasErr := 0
		for _, planDetail := range plan.PlanDetail {
			profileDetail := model.OwnerProfileDetail{}
			copier.Copy(&profileDetail, &planDetail)
			profileDetail.OwnerProfileId = profile.Id
			if _, err3 := db.MySQL().Insert(&profileDetail); err3 != nil {
				hasErr = 1
			}
		}

		if err2 == nil && affected > 0 && hasErr == 0 {
			return c.JSON(http.StatusOK, nil)
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func UpdatePlan(c echo.Context) (err error) {
	planId, err1 := strconv.Atoi(c.Param("id"))

	if err1 == nil {
		userId := c.Get("userId").(string)
		var plan = vm.RxPlan{}
		if err2 := c.Bind(&plan); err2 == nil {
			profile := model.OwnerProfile{}
			copier.Copy(&profile, &plan)
			//profile.Id = planId
			//profile.OwnerUserId = userId

			_, err3 := db.MySQL().Id(planId).Where("owner_user_id=?", userId).Update(&profile)
			_, err4 := db.MySQL().Delete(&model.OwnerProfileDetail{OwnerProfileId: planId})

			hasErr := 0
			for _, planDetail := range plan.PlanDetail {
				profileDetail := model.OwnerProfileDetail{}
				copier.Copy(&profileDetail, &planDetail)
				profileDetail.OwnerProfileId = planId
				if _, err4 := db.MySQL().Insert(&profileDetail); err4 != nil {
					hasErr = 1
				}
			}

			if err3 == nil && err4 == nil && hasErr == 0 {
				return c.JSON(http.StatusOK, nil)
			}
		}

	}

	return c.JSON(http.StatusBadRequest, nil)
}

func RenamePlan(c echo.Context) (err error) {
	planId, err1 := strconv.Atoi(c.Param("id"))

	if err1 == nil {
		userId := c.Get("userId").(string)
		var plan = vm.RxPlanRename{}
		if err2 := c.Bind(&plan); err2 == nil {
			profile := model.OwnerProfile{Pname: plan.Pname}
			affected, err3 := db.MySQL().Id(planId).Where("owner_user_id=?", userId).Cols("pname").Update(&profile)
			if err3 == nil && affected > 0 {
				return c.JSON(http.StatusOK, nil)
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func DeletePlan(c echo.Context) (err error) {
	planId, err1 := strconv.Atoi(c.Param("id"))

	if err1 == nil {
		userId := c.Get("userId").(string)
		//lockers := make([]model.Locker, 0)
		total, err2 := db.MySQL().Where("owner_profile_id = ?", planId).And("active = 1").Count(&model.Locker{})
		if err2 == nil {
			if total == 0 /*无锁正在使用此方案*/ {
				/*删除分段计费条款*/
				has, err3 := db.MySQL().Id(planId).Where("owner_user_id=?", userId).Get(&model.OwnerProfile{}) /*判断是本人的地锁条款*/
				if err3 == nil && has {
					_, err4 := db.MySQL().Where("owner_profile_id=?", planId).Delete(&model.OwnerProfileDetail{})
					if err4 == nil {
						/*删除主条款*/
						affected, err5 := db.MySQL().Id(planId).Delete(&model.OwnerProfile{})

						if err5 == nil && affected > 0 {
							return c.JSON(http.StatusOK, nil)
						}
					}
				}
			} else {
				return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "此方案尚在使用中"})
			}
		}
	}
	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}

func ListLocksUnderPlan(c echo.Context) (err error) {
	planId, err1 := strconv.Atoi(c.Param("id"))
	if err1 == nil {
		userId := c.Get("userId").(string)

		lockers := make([]model.Locker, 0)

		err2 := db.MySQL().Where("owner_user_id=?", userId).And("active = 1").And("owner_profile_id = ?", planId).Find(&lockers)

		if err2 == nil {
			txLocks := make([]vm.TxLockForList, 0)
			copier.Copy(&txLocks, lockers)
			return c.JSON(http.StatusOK, txLocks)
		}

	}

	return c.JSON(http.StatusBadRequest, &vm.TxInfo{Info: "未知错误"})
}
