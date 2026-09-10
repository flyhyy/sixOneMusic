import http from "@/http"


const AudioApi = {


    audioPath: (id: number) => `/api/audio/play/${id}`,
    audioLrc: (id: number) => http.get<string>(`/audio/play/lrc/${id}`, undefined, {
        responseType: "text",
    })

}

export default AudioApi