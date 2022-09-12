package constanta

const ActionInsert = "insert"
const ActionUpdate = "update"
const ActionGet = "get"
const ActionDelete = "delete"

const ActionIDInsert = 1
const ActionIDUpdate = 2
const ActionIDGet = 3
const ActionIDDelete = 4

// ---------------------------------- DB Table Name Constanta --------------------------------------------------------
const TableUsers = "users"
const TableResources = "resources"
const TableUserActivation = "user_activation"

// ---------------------------------- API Constanta --------------------------------------------------------
const APIResource = "/okami/auth/resource"
const APIUser = "/okami/auth/user"
const APITokenResource = "/okami/auth/token/resource"
const APIEmailValidate = "/okami/email/validate"
const APIEmailResend = "/okami/email/resend"
const APIPKCEStep1 = "okami/auth/pkce/step1"
const APIPKCEStep2 = "okami/auth/pkce/step2"
const APIPKCEStep3 = "okami/auth/pkce/step3"
