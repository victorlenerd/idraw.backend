(function () {
    
    function main() {
        const baseURL = "https://unilagpastquestions-com.appspot.com/notes/test-note-uuid";
        const imageURL = (fileName) => `https://storage.googleapis.com/idraw-app-images/${fileName}`;
        const root = document.getElementById("root");

        setInterval(() => {
            fetch(baseURL)
                .then((res) => res.json())
                .then((noteImages) => {
                    let maxVersion = Number.MIN_SAFE_INTEGER;
                    let mostRecentImageUrl = null;

                    noteImages.forEach((noteImage) => {
                        if (noteImage.version > maxVersion) {
                            maxVersion = noteImage.version
                            mostRecentImageUrl = noteImage.file_name;
                        }
                    });

                    const image = new Image();
                    image.src = imageURL(mostRecentImageUrl);

                    root.innerHTML = "";

                    root.appendChild(image);
                })
        }, 1000)
    }

    window.addEventListener('DOMContentLoaded', () => {
        console.log('DOM fully loaded and parsed');
        main()
    })

})();