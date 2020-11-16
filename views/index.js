(function () {
    
    function main() {
        const baseURL = "https://unilagpastquestions-com.appspot.com/notes/test-note-uuid";
        const imageURL = (fileName) => `https://storage.googleapis.com/idraw-app-images/${fileName}`;

        fetch(baseURL)
            .then((res) => res.json())
            .then((noteImages) => {
                console.log({ noteImages })
            })
    }

    window.addEventListener('DOMContentLoaded', () => {
        console.log('DOM fully loaded and parsed');
        main()
    })

})();