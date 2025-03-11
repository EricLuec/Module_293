function addPost() {
    let title = document.getElementById('title').value;
    let content = document.getElementById('content').value;
    let category = document.getElementById('category').value;

    if (title && content) {
        let posts = JSON.parse(localStorage.getItem('posts')) || [];
        let newPost = { title, content, category };
        posts.unshift(newPost);
        localStorage.setItem('posts', JSON.stringify(posts));

        document.getElementById('title').value = '';
        document.getElementById('content').value = '';

        alert('Beitrag erfolgreich erstellt!');
        window.location.href = 'index.html';
    } else {
        alert('Bitte alle Felder ausfüllen');
    }
}

function loadPosts() {
    let posts = JSON.parse(localStorage.getItem('posts')) || [];
    let postsContainer = document.getElementById('posts');
    postsContainer.innerHTML = '';

    posts.forEach(post => {
        let postElement = document.createElement('div');
        postElement.classList.add('post');
        postElement.setAttribute('data-category', post.category);
        postElement.innerHTML = `
            <h3>${post.title}</h3>
            <p>${post.content}</p>
            <p class='category'>Kategorie: ${post.category}</p>
            <button class='delete-btn' onclick='deletePost("${post.title}")'>Löschen</button>
        `;
        postsContainer.appendChild(postElement);
    });
}

function deletePost(title) {
    let posts = JSON.parse(localStorage.getItem('posts')) || [];
    posts = posts.filter(post => post.title !== title);
    localStorage.setItem('posts', JSON.stringify(posts));
    loadPosts();
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
