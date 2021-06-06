package storage

import (
	"context"
	"fmt"
)

//Photos структура для описания массива объектов таблицы photography
type Photos struct {
	PhotoID int
	ImageURL string
	IsImageMain bool
}

//Photography структура описывающая таблицу photography в БД
type Photography struct {
	ID    int
	Score string
	Photos map[string]Photos
}

//Select - метод для выборки данных из бд
func (p *Photography) Select(id int) Photography {

	sql := fmt.Sprintf(`SELECT json_build_object(
		'id',id,
		'score',score,
		'photos',(
				SELECT json_agg(json_build_object(
						'photo_id', Photo_image.photo_id,
						'image_url', Photo_image.image_url,
						'is_image_main', Photo_image.is_image_main
			))
			from Photo_image
			where Photography.id = Photo_image.photo_id
		)
	)
	FROM Photography
	WHERE id = %d`, id)

	photography := Photography{}

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}


	for rows.Next() {
		rows.Scan(&photography.ID, &photography.Score, &photography.Photos)
		if err != nil {
			fmt.Println(err)
		}
	}

	return photography
}