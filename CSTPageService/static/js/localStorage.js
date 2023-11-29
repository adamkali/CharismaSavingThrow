/// getter for localstorage data by key
/// @param {string} key
/// @return {string} or {null}
function getLocalStorage(key) {
    return localStorage.getItem(key);
}

/// setter for localstorage data by key
/// @param {string} key
/// @param {object} value
/// @return {void}
/// @throws {Error} if key is not a string
function setLocalStorage(key, value) {
    localStorage.setItem(key, value);
}
