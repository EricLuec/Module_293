async function loadPosts() {
    const res = await fetch("/api/posts");
    const posts = await res.json();
    const postsDiv = document.getElementById("posts");

    postsDiv.innerHTML = ""; // Posts leeren

    posts.forEach(post => {
        const div = document.createElement("div");
        div.className = "post";
        div.innerHTML = `
            <h3>${post.title}</h3>
            <p>${post.content}</p>
            <p><strong>Kategorie:</strong> ${post.category}</p>
            <button class="edit" onclick="editPost('${post.id}', '${post.title}', '${post.content}', '${post.category}')">Bearbeiten</button>
            <button class="delete" onclick="deletePost('${post.id}')">Löschen</button>
            <div id="comments-${post.id}" class="comments-container"></div>
            <div class="comment-form">
                <textarea id="commentContent-${post.id}" placeholder="Kommentar hinzufügen"></textarea>
                <button onclick="createComment('${post.id}')">Kommentar hinzufügen</button>
            </div>
        `;
        postsDiv.appendChild(div);

        loadComments(post.id);
    });
}

async function createPost(event) {
    event.preventDefault();
    const title = document.getElementById("title").value;
    const content = document.getElementById("content").value;
    const category = document.getElementById("category").value;

    const newPost = { title, content, category };

    const res = await fetch("/api/posts", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(newPost),
    });

    if (res.ok) {
        alert("Post erfolgreich erstellt!");
        loadPosts();
    }
}

async function editPost(id, title, content, category) {
    const newTitle = prompt("Neuer Titel:", title);
    const newContent = prompt("Neuer Inhalt:", content);
    const newCategory = prompt("Neue Kategorie:", category);

    if (newTitle && newContent && newCategory) {
        const updatedPost = {
            title: newTitle,
            content: newContent,
            category: newCategory,
        };

        const res = await fetch(`/api/posts/${id}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(updatedPost),
        });

        if (res.ok) {
            alert("Post erfolgreich aktualisiert!");
            loadPosts();
        }
    }
}

async function deletePost(id) {
    const confirmDelete = confirm("Möchten Sie diesen Post wirklich löschen?");
    if (confirmDelete) {
        const res = await fetch(`/api/posts/${id}`, {
            method: "DELETE",
        });

        if (res.ok) {
            alert("Post erfolgreich gelöscht!");
            loadPosts();
        }
    }
}

async function loadComments(postId) {
    const res = await fetch(`/api/comments/${postId}`);
    const comments = await res.json();
    const commentsDiv = document.getElementById(`comments-${postId}`);
    commentsDiv.innerHTML = "";

    comments.forEach(comment => {
        const commentDiv = document.createElement("div");
        commentDiv.innerHTML = `
            <p><strong>${comment.author}:</strong> ${comment.content}</p>
            <button onclick="editComment('${comment.id}', '${comment.content}')">Bearbeiten</button>
            <button onclick="deleteComment('${comment.id}')">Löschen</button>
        `;
        commentsDiv.appendChild(commentDiv);
    });
}

async function createComment(postId) {
    const content = document.getElementById(`commentContent-${postId}`).value;
    const author = prompt("Ihr Name:");

    if (content && author) {
        const newComment = { postId, content, author };

        const res = await fetch("/api/comments", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(newComment),
        });

        if (res.ok) {
            alert("Kommentar erfolgreich hinzugefügt!");
            loadComments(postId);
        }
    }
}

async function editComment(commentId, content) {
    const newContent = prompt("Neuer Kommentarinhalt:", content);
    if (newContent) {
        const updatedComment = { content: newContent };

        const res = await fetch(`/api/comments/${commentId}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(updatedComment),
        });

        if (res.ok) {
            alert("Kommentar erfolgreich aktualisiert!");
            loadPosts();
        }
    }
}

async function deleteComment(commentId) {
    const confirmDelete = confirm("Möchten Sie diesen Kommentar wirklich löschen?");
    if (confirmDelete) {
        const res = await fetch(`/api/comments/${commentId}`, {
            method: "DELETE",
        });

        if (res.ok) {
            alert("Kommentar erfolgreich gelöscht!");
            loadPosts();
        }
    }
}

document.getElementById("postForm").addEventListener("submit", createPost);
window.onload = loadPosts;
