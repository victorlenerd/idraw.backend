(function () {

    let intervalHandler = null;

    const baseURL = "https://idraw-app.df.r.appspot.com/notes";
    const imageURL = (fileName) => `https://storage.googleapis.com/idraw-images/${fileName}`;
    const root = document.getElementById("root");

    function renderLatestNoteImage(noteImages) {
        if (noteImages.length > 0) {
            let maxVersion = Number.MIN_SAFE_INTEGER;
            let mostRecentImageUrl = null;

            noteImages.forEach((noteImage) => {
                if (noteImage.version > maxVersion) {
                    maxVersion = noteImage.version;
                    mostRecentImageUrl = noteImage.file_name;
                }
            });

            const image = new Image();
            image.src = imageURL(mostRecentImageUrl);

            root.innerHTML = "";
            root.appendChild(image);
        } else {
            clearInterval(intervalHandler)
            // TODO: Empty note images array
        }
    }

    function fetchLatestNoteImage(noteUUID) {
        return () => {
            fetch(`${baseURL}/${noteUUID}`)
                .then((res) => res.json())
                .then(renderLatestNoteImage)
                .catch(() => {
                    // TODO: Handle error
                    clearInterval(intervalHandler)
                })
        };
    }

    function main() {
        const queryString = window.location.search;
        const noteUUIDQueryStrReg = /noteUUID=([^&]*)/;
        const matches = queryString.match(noteUUIDQueryStrReg)

        if (matches.length >= 2) {
            const noteUUID = matches[1];
            intervalHandler = setInterval(fetchLatestNoteImage(noteUUID), 1000)
        } else {
            // TODO: Handle missing note uuid query string
        }
    }

    window.addEventListener('DOMContentLoaded', main);
})();