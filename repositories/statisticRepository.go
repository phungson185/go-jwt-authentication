package repositories

import (
	"jwt-authen/database"
	"jwt-authen/models"
)

type StatisticRepo struct{}

func (s *StatisticRepo) UserRegisterInADay() (interface{}, error) {

	var total int64

	if err := database.Db.Model(&models.User{}).Where("date(created_at) = CURRENT_DATE").Find(&total).Error; err != nil {
		return nil, err
	}

	return total, nil
}

func (s *StatisticRepo) ListNewestItem() (interface{}, error) {

	var items []models.Item

	if err := database.Db.Model(&models.Item{}).Order("created_at desc").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (s *StatisticRepo) ListNewestAuction() (interface{}, error) {

	var auctions []models.Auction

	if err := database.Db.Model(&models.Auction{}).Order("created_at desc").Find(&auctions).Error; err != nil {
		return nil, err
	}

	return auctions, nil
}

func (s *StatisticRepo) SellestItem() (interface{}, error) {

	var item models.Item

	if err := database.Db.Raw("select * from items where id in (select item_id from transactions group by item_id order by count(item_id) desc limit 1)").Scan(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}
