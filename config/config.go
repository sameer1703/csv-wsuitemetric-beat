// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period               time.Duration `config:"period"`
	StateFileStorageType string        `config:"state_file_storage_type"`
	StateFileName        string        `config:"state_file_name"`
	StateFilePath        string        `config:"state_file_path"`
	AwsAccessKey         string        `config:"aws_access_key"`
	AwsSecretAccessKey   string        `config:"aws_secret_access_key"`
	AwsS3BucketName      string        `config:"aws_s3_bucket_name"`
	AwsRegion            string        `config:"aws_region"`
	FilesPrefix          string        `config:"file_prefix"`
	TimestampTemplate    string        `config:"timestamp_template"`
	TimestampFormat      string        `config:"timestamp_format"`
	EventTypeColumn      string        `config:"event_type_column"`
	BatchSize            int64         `config:"batch_size"`
}

var DefaultConfig = Config{
	Period:               1 * time.Second,
	StateFileStorageType: "disk",
	StateFileName:        "csvbeat",
	StateFilePath:        "/etc/csvbeat/",
}
