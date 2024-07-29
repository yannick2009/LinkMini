export const urlRegex: RegExp = new RegExp(
    '^(https?:\\/\\/)?' + // protocol
    '((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|' + // domain name
    '((\\d{1,3}\\.){3}\\d{1,3}))' + // OR ip (v4) address
    '(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*' + // port and path
    '(\\?[;&a-z\\d%_.~+=-]*)?' + // query string
    '(\\#[-a-z\\d_]*)?$',
    'i'
);

export const shortUrlRegex: RegExp = new RegExp(/^(http:\/\/|https:\/\/)?(localhost:[0-9]{1,5}|[a-zA-Z0-9.-]+)\/[a-zA-Z0-9]{7}$/gm);
