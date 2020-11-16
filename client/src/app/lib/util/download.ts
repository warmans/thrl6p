export function downloadPatch(filename: string, data: string) {
  filename = `${filename}.thrl6p`;
  const blob = new Blob([data], {type: 'application/json'});
  if (window.navigator.msSaveOrOpenBlob) {
    window.navigator.msSaveBlob(blob, filename);
  }
  else{
    const elem = window.document.createElement('a');
    elem.href = window.URL.createObjectURL(blob);
    elem.download = filename;
    document.body.appendChild(elem);
    elem.click();
    document.body.removeChild(elem);
  }
}

// create a valid file name e.g. USER PRESET 1 becomes user-preset-1
export function toFileName(name: string): string {
  return name
    .toLowerCase()
    .replace(/\s+/g, '-')
    .replace(/[^a-z0-9\-]+/g, '');
}
