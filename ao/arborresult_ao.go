package ao

import (
	"BrainTellServer/do"
	"BrainTellServer/models"
	"BrainTellServer/utils"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type UpdateArboResultAo struct {
	SomaId     int                    `json:"somaid"`
	SomaType   int                    `json:"somatype"`
	Insertlist []*models.TArborresult `json:"insertlist"`
	Deletelist []int                  `json:"deletelist"`
	Owner      string                 `json:"owner"`
	Image      string                 `json:"image"`
}

func UpdateArborResult(pa *UpdateArboResultAo) error {
	if _, err := utils.GetLocationTTL(fmt.Sprintf("Arbor_%d", pa.SomaId)); err != nil {
		return err
	}
	users, err := utils.GetLocationValue(fmt.Sprintf("Arbor_%d", pa.SomaId))
	if err != nil {
		return err
	}

	for _, user := range users {
		if user != pa.Owner {
			return errors.New(fmt.Sprintf("Not right user,%s %s\n", user, pa.Owner))
		}
	}

	if f, err := utils.In(users, pa.Owner); err != nil || !f {
		return errors.New(fmt.Sprintf("Not right user,%v %s\n", users, pa.Owner))
	}

	if utils.SetKeyTTL(fmt.Sprintf("Arbor_%d", pa.SomaId), 10*60) != nil {
		return err
	}

	_, err = do.UpdateArbor(&models.TArbor{
		Id:     pa.SomaId,
		Owner:  pa.Owner,
		Status: pa.SomaType,
	})

	if err != nil {
		return err
	}

	if len(pa.Deletelist) == 0 && len(pa.Insertlist) == 0 {
		return nil
	}
	//delete
	if len(pa.Deletelist) != 0 {
		_, err := do.DeleteArborResult(pa.Deletelist, pa.Owner)
		if err != nil {
			log.WithFields(log.Fields{
				"event":  "UpdateSomaList",
				"action": "delete",
			}).Warn(err)
			return err
		}
	}

	if len(pa.Insertlist) == 0 {
		return nil
	}
	_, err = do.InsertArborResult(pa.Insertlist)
	if err != nil {
		log.WithFields(log.Fields{
			"event":  "UpdateSomaList",
			"action": "InsertSomas",
		}).Warn(err)
		return err
	}
	return nil
}