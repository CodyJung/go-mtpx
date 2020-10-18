package main

type MtpDetectFailedError struct {
	error
}

type ConfigureError struct {
	error
}

type DeviceInfoError struct {
	error
}

type StorageInfoError struct {
	error
}

type NoStorageError struct {
	error
}

type ListDirectoryError struct {
	error
}

type FileNotFoundError struct {
	error
}

type InvalidPathError struct {
	error
}

type FileObjectError struct {
	error
}

type SendObjectError struct {
	error
}
