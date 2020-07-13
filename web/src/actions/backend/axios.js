import axios from 'axios';

// declare a request interceptor
axios.interceptors.request.use(config => {
    return config;
}, error => {
    return Promise.reject(error);
});

export const getAllCharacters = async () => {
    return axios({
            url: "http://localhost:9001/api/characters",
            method: "get",
            headers: {
                "Origin": "memequotes_front",
                "Accept": "application/json"
            }
        }
    );
};

export const getAllPhrasesForCharacter = async (characterId) => {
    return axios({
        url: `http://localhost:9001/api/character/${characterId}/phrases`,
        method: "get",
        headers: {
            "Origin": "memequotes_front",
            "Accept": "application/json"
        }
    })
}