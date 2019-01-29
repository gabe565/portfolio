'use strict'

export default function (url) {
    return new Promise((resolve, reject) => {
        let img = new Image()
        img.onload = () => {
            resolve(url)
        }
        img.onerror = () => {
            reject(url)
        }
        img.src = url
    })
}
