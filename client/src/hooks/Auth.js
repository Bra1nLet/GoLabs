import { useEffect, useState } from 'react';
import {computeHeadingLevel} from "@testing-library/react";

export function SetToken(token) {
    sessionStorage.setItem('token', 'Bearer %s'.replace("%s", JSON.stringify(token)).replaceAll("\"",""));
}

export function ExitFromAccount() {
    sessionStorage.setItem('token', '');
    window.location.reload(false);
}

export function Auth() {
    const isValid = ValidateToken(GetToken());

    if (!isValid){
        //sessionStorage.setItem('token', '');
        return false;
    }
    else if (isValid){
        return true;
    }
}

export function GetToken() {
    if (sessionStorage.getItem('token') === null){
        return ""
    }
    return sessionStorage.getItem('token');
}


export function ValidateToken() {
    const [data, setData] = useState();

    useEffect(() => {
        fetch('/auth/validate-token', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': GetToken()
            },
        })
            .then(response => response.json())
            .then(data => setData(data.isValid))
    }, []);

    console.log(data);
    if (data === "true") {
        return true
    }
    else if (data === "false"){
        return false
    }
}

export function GetUser(){
    const token = GetToken().replaceAll("\"", "");
    const [data, setData] = useState(false);

    useEffect(() => {
        fetch('/get-user', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
        })
            .then(response => {
                if (response !== undefined){
                setData(response.json())}
            })
            .catch(setData(false))
    }, []);

    if (data === false){
        return false;
    }
    else if (data !== false){
        return false;
    }
    return data;

}

