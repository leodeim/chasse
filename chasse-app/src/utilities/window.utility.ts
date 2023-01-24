enum Position {
    h = 'horizontal',
    v = 'vertical',
};

export type WindowProperties = {
    minDimension: number,
    position: Position
};

function getWindowProperties(): WindowProperties {
    const { innerWidth: width, innerHeight: height } = window;

    return {
        minDimension: Math.min(width, height),
        position: (width > height) ? Position.h : Position.v
    };
};

declare global {
    interface Window {
        clipboardData: any;
    }
  }

export function copyToClipboard(text) {
    if (window.clipboardData && window.clipboardData.setData) {
        return window.clipboardData.setData(`Text`, text);
    }
    else if (document.queryCommandSupported && document.queryCommandSupported(`copy`)) {
        let textarea = document.createElement(`textarea`);
        textarea.textContent = text;
        textarea.style.position = `fixed`;
        document.body.appendChild(textarea);
        textarea.select();
        try {
            return document.execCommand(`copy`);
        }
        catch (ex) {
            console.warn(`Copy to clipboard failed`, ex);
            return false;
        }
        finally {
            document.body.removeChild(textarea);
        }
    }
}

export { getWindowProperties, Position };