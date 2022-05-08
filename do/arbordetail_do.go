package do

import (
	"BrainTellServer/models"
	"BrainTellServer/utils"

	log "github.com/sirupsen/logrus"

	jsoniter "github.com/json-iterator/go"
)

type ArborDetail struct {
	Id      int     `json:"id"`
	ArborId int     `json:"arborId"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Z       float64 `json:"z"`
	Type    int     `json:"type"`
	Owner   string  `json:"owner"`
}

func InsetArborDetail(pa []*models.TArbordetail) (int, error) {
	jsonpa, _ := jsoniter.MarshalToString(pa)

	affected, err := utils.DB.NewSession().Insert(pa)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "Insert arbordetail failed",
			"pa":    jsonpa,
		}).Warnf("%v\n", err)
		return 0, err
	}
	log.WithFields(log.Fields{
		"event": "Insert arbordetail success",
		"pa":    jsonpa,
	}).Warnf("%v\n", err)
	return int(affected), nil
}

func DeleteArbordetail(pa []*models.TArbordetail) (int, error) {
	jsonpa, _ := jsoniter.MarshalToString(pa)
	affected, err := utils.DB.NewSession().Update(&models.TArbordetail{
		Isdeleted: 1,
	}, pa)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "delete arbordetail failed",
			"pa":    jsonpa,
		}).Warnf("%v\n", err)
		return 0, err
	}
	log.WithFields(log.Fields{
		"event": "delete arbordetail success",
		"pa":    jsonpa,
	}).Warnf("%v\n", err)
	return int(affected), nil

}

func QueryArborDetail(pa *models.TArbordetail) ([]*ArborDetail, error) {
	jsonpa, _ := jsoniter.MarshalToString(pa)
	res := make([]*ArborDetail, 0)
	rows := make([]*models.TArbordetail, 0)
	err := utils.DB.NewSession().Find(&rows, pa)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "query arbordetail failed",
			"pa":    jsonpa,
		}).Warnf("%v\n", err)
		return nil, err
	}
	for _, v := range rows {
		res = append(res, &ArborDetail{
			Id:      v.Id,
			ArborId: v.Arborid,
			X:       v.X,
			Y:       v.Y,
			Z:       v.Z,
			Type:    v.Type,
			Owner:   v.Owner,
		})
	}
	return res, nil
}
