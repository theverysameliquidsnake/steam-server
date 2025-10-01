export default {
    getCurrentTime() {
        const now = new Date();
        const hours = now.getHours();
        const minutes = now.getMinutes();
        const seconds = now.getSeconds();

        return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
    },

    calculatePagination(current, last) {
        let delta = 10;
        let left = current - delta;
        let right = current + delta + 1;
        let range = [];
        let rangeWithDots = [];
        let l;

        for (let i = 1; i <= last; i++) {
            if (i === 1 || i === last || i >= left && i < right) {
                range.push(i);
            }
        }

        for (let i of range) {
            if (l) {
                if (i - l === 2) {
                    rangeWithDots.push(l + 1);
                } else if (i - l !== 1) {
                    rangeWithDots.push("...");
                }
            }
            rangeWithDots.push(i);
            l = i;
        }

        return rangeWithDots;
    }
}