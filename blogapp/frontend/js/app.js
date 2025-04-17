const form = document.getElementById("postForm");
const postsDiv = document.getElementById("posts");

form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const title = document.getElementById("title").value.trim();
    const content = document.getElementById("content").value.trim();

    if (!title || !content) return;

    await fetch("/api/posts", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, content })
    });

    form.reset();
    loadPosts();
});

async function loadPosts() {
    const res = await fetch("/api/posts");
    const posts = await res.json();
    postsDiv.innerHTML = "";

    posts.forEach(post => {
        const div = document.createElement("div");
        div.className = "post";
        div.innerHTML = `
      <div class="actions">
        <button onclick="editPost('${post.id}', '${post.title}', \`${post.content.replace(/`/g, "\\`")}\`)">✏️</button>
        <button onclick="deletePost('${post.id}')">🗑️</button>
      </div>
      <h3>${post.title}</h3>
      <p>${post.content}</p>
    `;
        postsDiv.appendChild(div);
    });
}

async function deletePost(id) {
    await fetch(`/api/posts/${id}`, { method: "DELETE" });
    loadPosts();
}

function editPost(id, title, content) {
    const newTitle = prompt("Neuer Titel:", title);
    const newContent = prompt("Neuer Inhalt:", content);

    if (newTitle && newContent) {
        fetch(`/api/posts/${id}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ title: newTitle, content: newContent })
        }).then(loadPosts);
    }
}

loadPosts();
