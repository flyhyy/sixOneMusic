package repository

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"fnMusicServe/interal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// 文件夹名称
	rootDir, _ := os.Getwd()
	dbDir := "data"
	dbName := "nas_music.db"

	// 查找是否有文件夹
	_, err := os.Stat(dbDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// 文件不存在,就创建
			err := os.Mkdir(dbDir, os.ModePerm)
			if err != nil {
				log.Fatalf("文件创建失败:%v", err)
			}

		} else {
			// 文件发生其他错误
			log.Fatalf("发生错误%v", err)
		}
	}
	//  独立处理数据库连接（无论文件夹是刚建的还是已经存在的，都要走这里）
	// 拼接地址
	dbPath := filepath.Join(rootDir, dbDir, dbName)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败:%v", err)
	}

	log.Printf("数据库连接成功,路径：%v", dbPath)

	err = db.AutoMigrate(&model.UserDB{}, &model.PlayListDB{}, &model.SongDB{}, &model.FolderDB{}, &model.StyleDB{}, &model.SongCollectsDB{}, &model.SingerDB{})
	if err != nil {
		log.Fatalf("数据库初始化失败：%v", err)
	}
	log.Println("数据库表初始化成功")

	return db
}
