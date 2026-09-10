
import http from '@/http/index'
import type { LoginFormData, ResponseLogin } from '@/pages/login/types'
const LoginApi = {

    login: (data: LoginFormData) => {
        return http.post<ResponseLogin>("/auth/login", data)
    }

}

export default LoginApi