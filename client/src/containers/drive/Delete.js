import React from 'react';
import { useRef, useState } from 'react';
import {Authorize}  from "../../hooks/";
import {ExitFromAccount, GetToken} from "../../hooks/Auth";
import './Drive.css'

const DeleteDrive = () => {
    const path = useRef("");
    const nameFile = useRef("");

    const handleButtonCreate = () => {
        const formData = new FormData();
        formData.append('path', path.current.value);
        formData.append('filename', nameFile.current.value);
        const authOptions = {
            method: "POST",
            body: formData,
            headers: {
                'Authorization': GetToken()
            },
        };

        fetch('/drive/delete', authOptions);
        //window.location.reload(false);
    };


    return (
        <div>
            <h1>Delete</h1>
            <label htmlFor="path">Path</label>
            <br />
            <input ref={path} name="path" type="text" />
            <br />
            <label htmlFor="name">Name</label>
            <br />
            <input ref={nameFile} name="name" type="text" />
            <br/>
            <br/>
            <button onClick={handleButtonCreate}>Delete</button>
        </div>
    );
};
export default DeleteDrive;