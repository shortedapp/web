export const findMinMax = arr => {
  let min = arr[0].y
  let max = arr[0].y
  for (let i = 1, len = arr.length; i < len; i++) {
    let v = arr[i].y
    if (v) {
      min = v < min ? v : min
      max = v > max ? v : max
    }
  }
  return [min, max]
}

Date.prototype.getMonthName = function (lang) {
  lang = lang && lang in Date.locale ? lang : 'en'
  return Date.locale[lang].month_names[this.getMonth()]
}

Date.prototype.getMonthNameShort = function (lang) {
  lang = lang && lang in Date.locale ? lang : 'en'
  return Date.locale[lang].month_names_short[this.getMonth()]
}

Date.locale = {
  en: {
    month_names: [
      'January',
      'February',
      'March',
      'April',
      'May',
      'June',
      'July',
      'August',
      'September',
      'October',
      'November',
      'December'
    ],
    month_names_short: [
      'Jan',
      'Feb',
      'Mar',
      'Apr',
      'May',
      'Jun',
      'Jul',
      'Aug',
      'Sep',
      'Oct',
      'Nov',
      'Dec'
    ]
  }
}

export const getMonthlyName = (t) => {
    var dateParts = t.split("/");
    const date = new Date(dateParts[2], dateParts[1] - 1, dateParts[0])
    // console.log('utils:date:',date)
    // console.log('utils:getMonthName:',date.getMonthName())
    return date.getMonthName()
}
export const getYearlyName = (t) => {
    var dateParts = t.split("/");
    const date = new Date(dateParts[2], dateParts[1] - 1, dateParts[0])
    // console.log('utils:date:',date)
    // console.log('utils:getMonthName:',date.getMonthName())
    return date.getMonthName()
}
export const getYearlyNameShort = (t) => {
    var dateParts = t.split("/");
    const date = new Date(dateParts[2], dateParts[1] - 1, dateParts[0])
    // console.log('utils:date:',date)
    // console.log('utils:getMonthName:',date.getMonthName())
    return `${date.getMonthNameShort()} ${date.getFullYear()}`
}
export const getMonthlyNameShort = (t) => {
    // console.log('utils:t:',t)
    var dateParts = t.split("/");
    const date = new Date(dateParts[2], dateParts[1] - 1, dateParts[0])
    // const date = new Date(t)
    // console.log('utils:date:',date)
    // console.log('utils:getMonthNameShort:',date.getMonthName())
    return date.getMonthNameShort()
}


/**
 * getLinks
 * @param {*} metadata
 * parses metadata generates a number of external links made available on profile/stock legend views
 * Bloomberg: ASN -> https://www.bloomberg.com/quote/ASN:AU 
 * Google Finance: ASN -> https://www.google.com/search?q=ASX:ASN
 * Yahoo FInance: ASN -> https://finance.yahoo.com/quote/ASN.AX?ltr=1
 * ASX: ASN -> https://www.asx.com.au/asx/share-price-research/company/ASN
 * HotCopper : ASN -> https://hotcopper.com.au/asx/asn/
 * Twitter (from metadata)
 * Company website (from metadata)
 * 
 */
export const getLinks = (metadata) => {
    return [
        {
            name: 'company',
            url: metadata.company_url
        },
        {
            name: 'twitter',
            url: metadata.twitter_url,
        },
        {
            name: 'asx',
            url: `https://www.asx.com.au/asx/share-price-research/company/${metadata.code}`
        },
        {
            name: 'hotcopper',
            url: `https://hotcopper.com.au/asx/${metadata.code}/`
        },
        {
            name: 'bloomberg',
            url: `https://www.bloomberg.com/quote/${metadata.code}:AU`
        },
        {
            name: 'yahoo',
            url: `https://finance.yahoo.com/quote/${metadata.code}.AX`
        },
        {
            name: 'google',
            url: `https://www.google.com/search?q=ASX:${metadata.code}`
        }

    ]

}