package models

type ResponseTimeLog struct {
	Id         uint `gorm:"PrimaryKey;AutoIncrement"`
	Request    string
	TimeDb     float32
	TimeRedis  float32
	TimeMemory float32
}

func LogTracker(time float32, request string, storage string) {
	switch storage {
	case "db":
		result := DB.Model(&ResponseTimeLog{}).Where("request = ?", request).Update("time_db", time)
		if result.RowsAffected == 0 {
			newRec := ResponseTimeLog{Request: request, TimeDb: time, TimeRedis: 0, TimeMemory: 0}
			DB.Create(&newRec)
		}
	case "redis":
		DB.Model(&ResponseTimeLog{}).Where("request = ?", request).Update("time_redis", time)
	case "memory":
		DB.Model(&ResponseTimeLog{}).Where("request = ?", request).Update("time_memory", time)
	}
}
