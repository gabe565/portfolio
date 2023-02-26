export default (url) => {
  return new Promise((resolve, reject) => {
    const img = new Image();
    img.onload = () => resolve(url);
    img.onerror = () => reject(Error(`Failed to load ${url}`));
    img.src = url;
  });
};
