import React, { useRef, useState, useEffect } from 'react';
import { SetToken } from "../../hooks/Auth";

const Login = () => {
    const username = useRef(null);
    const password = useRef(null);
    const [AuthDetail, setAuthDetail] = useState();


    const handleButtonClick = () => {
        const authOptions = {
            method: 'POST',
            headers: {
                Accept: 'application/json',
            },
            body: JSON.stringify({
                username: username.current.value,
                password: password.current.value,
            }),
        };

        fetch('/auth/new-token', authOptions)
         .then(response => response.json())
         .then(response => login(response))
         .catch(error => {
            console.error('Error:', error);
         });
    };

    function login(data) {
        SetToken(data["token"]);
        window.location.reload(false);
    }

    return (
        <div>
            <h1>Login</h1>
            <label htmlFor="user">Username</label>
            <br />
            <input ref={username} name="user" type="text" />
            <br />
            <label htmlFor="pass">Password</label>
            <br />
            <input ref={password} name="pass" type="password" />
            <br />
            <button onClick={handleButtonClick}>Login</button>
        </div>
    );
};

export default Login;
