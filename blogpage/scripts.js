function addPost() {
    let title = document.getElementById('title').value;
    let content = document.getElementById('content').value;
    let category = document.getElementById('category').value;

    if (title && content) {
        let post = document.createElement('div');
        post.classList.add('post');
        post.innerHTML = `
            <h3>${title}</h3>
            <p>${content}</p>
            <p class='category'>Kategorie: ${category}</p>
            <button class='delete-btn' onclick='deletePost(this)'>Löschen</button>
        `;
        post.setAttribute('data-category', category);
        document.getElementById('posts').prepend(post);

        document.getElementById('title').value = '';
        document.getElementById('content').value = '';
    } else {
        alert('Bitte alle Felder ausfüllen');
    }
}

function deletePost(button) {
    button.parentElement.remove();
}

function filterPosts() {
    let selectedCategory = document.getElementById('filter').value;
    let posts = document.querySelectorAll('.post');

    posts.forEach(post => {
        if (selectedCategory === 'Alle' || post.getAttribute('data-category') === selectedCategory) {
            post.style.display = 'block';
        } else {
            post.style.display = 'none';
        }
    });
}
