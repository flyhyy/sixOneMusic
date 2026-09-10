export interface LoginFormData {
    username: string
    password: string
}

export interface RegisterFormData extends LoginFormData {
    confirmPassword: string
}

export interface ResponseLogin {
    token: string
}