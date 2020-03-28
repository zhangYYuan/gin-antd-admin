const moment = require('moment')
const CryptoJS = require('crypto-js')

export const encryption = (word) => {
  const parse = CryptoJS.enc.Utf8.parse(word)
  const date = `${moment().format("YYYY-MM-DD")}>musee`
  const key = CryptoJS.enc.Utf8.parse(date)
  const iv = key
  const encrypted = CryptoJS.AES.encrypt(parse, key, { iv, mode: CryptoJS.mode.CBC, padding: CryptoJS.pad.Pkcs7 })
  const result = encrypted.toString()
  return result
}
