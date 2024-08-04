/* -------------------------------- HELPERS ----------------------------------*/

// Checking rather an object is empty
export const isEmptyObject = (obj) => {
  return Object.keys(obj).length === 0 && obj.constructor === Object;
};

// Checking rather an array is empty
export const isEmptyArray = (arr) => {
  return Array.isArray(arr) && arr.length === 0;
};
