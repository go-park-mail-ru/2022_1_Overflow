package models

type SignInForm struct {
	Username string `json:"username" validate:"nonzero,max=45,regexp=^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"`
	Password string `json:"password" validate:"nonzero,max=45"`
}

type SignUpForm struct {
	Firstname string `json:"first_name" validate:"nonzero,max=45"`
	Lastname string `json:"last_name" validate:"nonzero,max=45"`
	Username string `json:"username" validate:"nonzero,max=45,regexp=^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"`
	Password string `json:"password" validate:"nonzero,max=45"`
	PasswordConfirmation string `json:"password_confirmation" validate:"nonzero,max=45"`
}

type MailForm struct {
	Addressee string `json:"addressee" validate:"nonzero,max=45"`
	Theme     string `json:"theme" validate:"nonzero,max=45"`
	Text      string `json:"text" validate:"nonzero"`
	Files     string `json:"files"`
}

type ProfileSettingsForm struct {
	Firstname string `json:"first_name" validate:"nonzero,max=45"`
	Lastname string `json:"last_name" validate:"nonzero,max=45"`
}

type ChangePasswordForm struct {
	OldPassword string `json:"password_old" validate:"nonzero,max=45"`
	NewPassword string `json:"password_new" validate:"nonzero,max=45"`
	NewPasswordConf string `json:"password_new_confirmation" validate:"nonzero,max=45"`
}

type ReadMailForm struct {
	Id int32 `json:"id" validate:"nonzero"`
	IsRead bool `json:"isread"`
}

type DeleteMailForm struct {
	Id int32 `json:"id" validate:"nonzero"`
}

type AddFolderForm struct {
	FolderName string `json:"folder_name" validate:"nonzero"`
}

type AddMailToFolderByIdForm struct {
	FolderName string `json:"folder_name" validate:"nonzero"`
	MailId int32 `json:"mail_id" validate:"nonzero"`
	Move bool `json:"move"`
}

type AddMailToFolderByObjectForm struct {
	FolderName string `json:"folder_name" validate:"nonzero"`
	Mail MailForm `json:"form"`
}

type MoveFolderMailForm struct {
	FolderNameSrc string `json:"folder_name_src" validate:"nonzero"`
	FolderNameDest string `json:"folder_name_dest" validate:"nonzero"`
	MailId int32 `json:"mail_id" validate:"nonzero"`
}

type ChangeFolderForm struct {
	FolderName string `json:"folder_name" validate:"nonzero"`
	NewFolderName string `json:"new_folder_name" validate:"nonzero"`
}

type DeleteFolderForm struct {
	FolderName string `json:"folder_name" validate:"nonzero"`
}

type DeleteFolderMailForm struct {
	FolderName string `json:"folder_name" validate:"nonzero"`
	MailId int32 `json:"mail_id" validate:"nonzero"`
	Restore bool `json:"restore"`
}