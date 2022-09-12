package constanta

const PathEnv = "okami.auth.backend/config/"

//------------------------- Header Request -----------------------------
const X_REQUEST_ID = "X-REQUEST-Id"
const X_DEVICE = "X-DEVICE"
const ContentTypeHTML = "text/html"
const HeaderKeyContentType = "Content-Type"
const HeaderValueContentType = "application/json"
const TokenHeaderNameConstanta = "Authorization"
const HeaderKeySecretCode = "Secret-Code"

//--------------------------- Default Language -------------------------
const DEFAULT_LANGUAGE = "id-ID"

//--------------------------- Default Time Format -------------------------
const DefaultTimeFormat = "2006-01-02 15:04:05"

// ---------------------------------- Context Name Constanta --------------------------------------------------------
const ApplicationContextConstanta = "application_context"
const TimestampSignatureHeaderNameConstanta = "X-Timestamp"
const SignatureHeaderNameConstanta = "X-Signature"

// ---------------------------------- Param Name Constanta --------------------------------------------------------
const EMAILTO_KEY = "email_to"
const USERNAME_KEY = "username"
const CODE_key = "code"
const CLIENTID_KEY = "clientID"

// ---------------------------------- Prefix Constanta --------------------------------------------------------
const PrefixDataDeleted = "_deleted_"
const PrefixLog = "::"
const PrefixPKCE = "-"

// ---------------------------------- General Constanta --------------------------------------------------------
const PayloadStatusCode = "OK"
const PayloadStatusSuccessTrue = true
const PayloadStatusSuccessFalse = false
const StringTrue = "true"
const StringFalse = "false"
const StringUp = "Up"
const StringFailed = "Failed"

const RequestPOST = "POST"
const RequestGET = "GET"
const RequestPUT = "PUT"
const RequestDELETE = "DELETE"

const FieldDataDeleted = "data_deleted"
const FieldDataNewest = "data_newest"

const EmailOkamiProject = "info.okami.project@gmail.com"
const EmailAppPassword = "dfdaohklehxnsccl" // https://support.google.com/mail/answer/185833?hl=en
const EmailHostGmail = "smtp.gmail.com"
const EmailHostGmailWithPort = "smtp.gmail.com:587"
const EmailTemplatePath = "D:\\okami-project\\asset\\template-email.html"
const EmailVerificationSubject = "Email Verification"
const EmailCodeMinimumLength = 1000000
const EmailCodeMaximumLength = 9999999
const EmailExpiredHour = 2 // masa email aktivasi selama 2 jam

const ActivationTemplatePath = "D:\\okami-project\\asset\\template-activation-status.html"

const ForbiddenResourceAuth = "auth" // hanya resource gate yang dapat access ke auth

const EncryptSHA256 = "SHA256"
const AuthTypePKCE = "pkce"

// ---------------------------------- User Constanta --------------------------------------------------------
const UserDeleted = 0
const UserPending = 1 // user belum melakukan aktivasi
const UserActive = 2

// ---------------------------------- Code Constanta --------------------------------------------------------
const CodeValidationFailed = "OKAMI-370001-VALIDATION-FAILED"
const CodeFieldIsEmpty = "OKAMI-370002-FIELD-ISEMPTY"
const CodeGetDataFailed = "OKAMI-370003-GETDATA-FAILED"
const CodeDBServerError = "OKAMI-370004-DBSERVER-ERROR"
const CodeRegistrationFailed = "OKAMI-370005-REGISTRATION-FAILED"
const CodeSendEmailFailed = "OKAMI-370006-SENDEMAIL-FAILED"
const CodeAuthorizationFailed = "OKAMI-370007-AUTHORIZATION-FAILED"
const CodeLoginFailed = "OKAMI-370008-LOGIN-FAILED"
const CodeGenerateFailed = "OKAMI-370009-GENERATE-FAILED"
const CodeVerifyFailed = "OKAMI-370009-VERIFY-FAILED"
const CodeRequestFailed = "OKAMI-370010-REQUEST-FAILED"

// ---------------------------------- Message Constanta --------------------------------------------------------

// REGEX
const RegexValidationFailed = "Regex Validation Failed"
const FormatEmailIsWrong = "Format Email Is Wrong"
const FormatUsernameIsWrong = "Format Username Is Wrong"
const FormatPasswordIsWrong = "Format Password Is Wrong"

// COMMAND OR WARNING GENERAL
const PleaseCheckYourEmail = "Please Check Your Email"
const RegistrationSuccess = "Registration Success !!"
const EmailResendFailed = "Failed to resend activation account"
const EmailResendSuccess = "Success to resend activation account"
const PleaseCheckYourParams = "Please Check Your Params"
const InvalidToken = "Invalid Token"
const DataDoesNotMatch = "Data Doesn't Match"
const ActivationFailed = "Failed to Activate Account"
const ActivationSuccess = "Account Activation Successful"
const PleaseCompleteYourProfile = "Please click the button below to complete your profile"
const MessageResendSuccess = "Your activation email has been sent" + " !! " + PleaseCheckYourEmail
const MessageRegistrationSuccess = RegistrationSuccess + " !! " + PleaseCompleteYourProfile
const InvalidPassword = "Invalid Password"
const LoginSuccess = "Login Success !!"
const GenerateSuccess = "Success to Generate !!"
const GenerateFailed = "Failed to Generate !!"
const VerifySuccess = "Success to Verify !!"
const VerifyFailed = "Failed to Verify !!"
const AccessForbidden = "Forbidden Access"
const DataNotFound = "Data Not Found !!"
const DataFounded = "Data Founded"
const PleaseCheckYourConnection = "Please check your connection"

const DeleteDataFailed = "Failed to delete data"
const DeleteDataSuccess = "Success to delete data"
const UpdateDataFailed = "Failed to update data"
const UpdateDataSuccess = "Success to update data"
const InsertDataFailed = "Failed to insert data"
const InsertDataSuccess = "Success to insert data"

const NoteActivationFailedUserNotFound = "Sorry, we can't process your activation ☹"
const NoteActivationSuccessUserActivated = "Yeay, your account has been activate ツ"
const NoteActivationFailedResendCounter = "Sorry, you have requested more than 3 times today ☹"

const AuthSuccess = "Authenticated Successfully"

// USER
const UserHasNotActivated = "User Has Not Activated"         // UserPending
const UserHasAlreadyActivated = "User Has Already Activated" // UserActive
const UserHasBlocked = "User Has Blocked"                    // UserDeleted
const UsernameOrEmailAlreadyExist = "Username Or Email Already Exist"
const UserNotFound = "User Not Found"

// BUTTON
const ClickHere = "Click Here !!"
const BackToHome = "Back To Home"
const ResendEmail = "Resend Email"
const LoginHere = "Login Here !!"
const LinkOkami = "https://google.com"
