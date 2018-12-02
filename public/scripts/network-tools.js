"use strict";

const jsonToQueryStr = data => {
    const item = [];
    Object.keys(data).map(key => item.push(`${key}=${encodeURIComponent(data[key])}`));
    return item.join("&");
};

const newToRemote = (data, path) => new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.timeout = 15 * 1000;
    xhr.ontimeout = () => reject(new Error('time is up!'));
    xhr.open("POST", path);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("cache-control", "no-cache");
    xhr.send(jsonToQueryStr(data));
    xhr.onreadystatechange = () => {
        if (4 === xhr.readyState) {
            const result = JSON.parse(xhr.responseText);
            if (200 === xhr.status) {
                if (200 !== result.code) reject(result.msg);
                resolve(result.data);
            } else reject(result.msg);
        }
    };
});

const appendToRemote = (data, path) => new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.timeout = 15 * 1000;
    xhr.ontimeout = () => reject(new Error('time is up!'));
    xhr.open("PUT", path);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("cache-control", "no-cache");
    xhr.send(jsonToQueryStr(data));
    xhr.onreadystatechange = () => {
        if (4 === xhr.readyState) {
            const result = JSON.parse(xhr.responseText);
            if (200 === xhr.status) {
                if (200 !== result.code) reject(result.msg);
                resolve(result.data);
            } else reject(result.msg);
        }
    };
});

const deleteFromRemote = (data, path) => new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.timeout = 15 * 1000;
    xhr.ontimeout = () => reject(new Error('time is up!'));
    xhr.open("DELETE", path+"?"+jsonToQueryStr(data));
    xhr.setRequestHeader("cache-control", "no-cache");
    xhr.send();
    xhr.onreadystatechange = () => {
        if (4 === xhr.readyState) {
            const result = JSON.parse(xhr.responseText);
            if (200 === xhr.status) {
                if (200 !== result.code) reject(result.msg);
                resolve(result.data);
            } else reject(result.msg);
        }
    };
});

const getListFromRemote = path => new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.timeout = 15 * 1000;
    xhr.ontimeout = () => reject(new Error('time is up!'));
    xhr.open("GET", path);
    xhr.send();
    xhr.onreadystatechange = () => {
        if (4 === xhr.readyState) {
            const result = JSON.parse(xhr.responseText);
            if (200 === xhr.status) {
                if (200 !== result.code) reject(result.msg);
                resolve(result.data);
            } else reject(result.msg);
        }
    };
});