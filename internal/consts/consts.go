package consts

const ACTIVE = "active"
const INACTIVE = "inactive"

// Notification Status
const SENT = "sent"
const SEEN = "seen"

// Features
const CREATE_JOB = "create job"
const UPDATE_JOB = "update job"
const DELETE_JOB = "delete job"
const CREATE_SCHEDULE = "create schedule"
const UPDATE_SCHEDULE = "update schedule"
const DELETE_SCHEDULE = "delete schedule"

// Response Message
const SUCCESS_CREATE = "Successfully created"
const SUCCESS_UPDATE = "Successfully updated"
const SUCCESS_DELETE = "Successfully deleted"
const FAILED_MJML = "Failed to Convert MJML To HTML: "
const ERROR_UNAUTHORIZED = "You do not have permission to perform this action. "

// Account Type
const USER = "user"
const COMPANY = "company"

// Token Type
const VALIDATE_ACCOUNT_TOKEN = "validate account"
const RESET_PASSWORD_TOKEN = "reset password"

// Application File Type
const RESUME = "resume"
const OTHER = "other"

// Context Variable Keys
const SESSION_DATA = "sessionData"
const TOKEN_CLAIM = "tokenClaim"
const TOKEN = "token"

// Environment
const DOMAIN = "DOMAIN"
const SUPPORT_EMAIL = "SUPPORT_EMAIL"
const SMTP_FROM = "SMTP_FROM"
const SMTP_HOST = "SMTP_HOST"
const SMTP_PORT = "SMTP_PORT"
const SMTP_PASSWORD = "SMTP_PASSWORD"
const SECRET = "SECRET"
const API_URL = "API_URL"
const AWS_BUCKET = "AWS_BUCKET"
const AWS_REGION = "AWS_REGION"
const AWS_ACCESS_KEY_ID = "AWS_ACCESS_KEY_ID"
const AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
const CERT_SERVICE_API_URL = "CERT_SERVICE_API_URL"

// Cache Prefix
const CACHE_USER_ID = "UserID:"
const CACHE_USER_PLAN_ID = "UserPlanID:"
const CACHE_COMPANY_ID = "CompanyID:"
const CACHE_COMPANY_ACCOUNTS_BY_COMPANY_ID = "CompanyAccountsByCompanyID:"
const CACHE_COMPANY_ACCOUNTS_BY_USER_ID = "CompanyAccountsByUserID:"
const CACHE_PERMISSIONS_BY_ROLE_ID = "PermissionsByRoleID:"
const CACHE_JOBS = "Jobs:"
const CACHE_JOBS_BY_COMPANY_ID = "JobsByCompanyID:"
const CACHE_JOB_ID = "JobID:"
const CACHE_APPLICATION_ID = "ApplicationID:"
const CACHE_APPLICATIONS_BY_JOB_ID = "ApplicationsByJobID:"
const CACHE_APPLICATIONS_BY_USER_ID = "ApplicationsByUserID:"
const CACHE_NOTIFICATION_ID = "NotificationID:"
const CACHE_NOTIFICATIONS_BY_ACCOUNT_ID = "NotificationsByAccountID:"
const CACHE_SUGGEST_SKILL_ID = "SuggestSkillID:"
const CACHE_FEATURE_ID = "FeatureID:"
