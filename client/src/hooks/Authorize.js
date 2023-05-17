import { useEffect, useState } from 'react';
import {SetToken} from './Auth';

export function Authorize(username, password) {
    try {
        const response = fetch('/auth/new-token',
            {
            method: 'POST',
            headers: {
                Accept: 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        });

        if (!response.ok) {
             new Error(`Error! status: ${response.status}`);
        }
        const result = response.json();

        if (result != null){
            SetToken(result["token"]);
        }
    } catch (err) {
        console.log(err.message);
    }
}

export default Authorize;