// Vite 版本：使用 import.meta.glob
export const getSvgNames = async () => {
    const svgFiles = import.meta.glob('@/assets/icons/*.svg', { eager: true });
    const svgNames = [];
    for (const path in svgFiles) {
      const fileName = path.split('/').pop().replace(/\.\w+$/, '');
      svgNames.push(fileName);
    }
    return svgNames;
};