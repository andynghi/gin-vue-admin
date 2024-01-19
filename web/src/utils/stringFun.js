/* eslint-disable */
export const toUpperCase = (str) => {
    if (str[0]) {
        return str.replace(str[0], str[0].toUpperCase())
    } else {
        return ''
    }
}

export const toLowerCase = (str) => {
    if (str[0]) {
        return str.replace(str[0], str[0].toLowerCase())
    } else {
        return ''
    }
}

// Convert camel case to underscore
export const toSQLLine = (str) => {
    if (str === 'ID') return 'ID'
    return str.replace(/([A-Z])/g, "_$1").toLowerCase();
}

// Convert underscore to camel case
export const toHump = (name) => {
    return name.replace(/\_(\w)/g, function(all, letter) {
        return letter.toUpperCase();
    });
}