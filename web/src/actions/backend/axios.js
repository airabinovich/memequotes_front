import axios from 'axios';

// declare a request interceptor
axios.interceptors.request.use(config => {
    return config;
}, error => {
    return Promise.reject(error);
});

export const getAllCharacters = async () => {
    let response = await axios({
            url: "http://localhost:9001/api/characters",
            method: "get",
            headers: {
                "Origin": "memequotes_front",
                "Accept": "application/json"
            }
        }
    );
    /*var data = []
    response.then(r => {
        console.log("RESPONSE: ", r)
        data = r.data
    })
        .catch(error => {
            console.log(error)
        });
    console.log("DATA:", data)*/
    return response;
};