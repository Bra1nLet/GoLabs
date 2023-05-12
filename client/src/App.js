import React, { useEffect, useState } from 'react';
import './App.css';
import { getList } from './list/list';

function Start(){
    var res = fetch('/auth/test')
    return "test"
}

function Star() {
    const requestOptions = {
        method: "POST",
        mode: "no-cors",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(this)
    }
    var response = fetch("/auth/test", requestOptions).then((data) => {
        return data.text()
    });
    //var result = response.json();
    return ("test");


}

export default Star;