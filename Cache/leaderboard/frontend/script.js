async function updateScore() {
    const user_id = document.getElementById('user_id').value;
    const score = document.getElementById('score').value;

    const response = await fetch('http://localhost:8080/api/scores', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ user_id, score })
    });

    if (response.ok) {
        alert('Score updated!');
        loadLeaderboard();
    }
}

async function loadLeaderboard() {
    const response = await fetch('http://localhost:8080/api/leaderboard?top=10');
    const data = await response.json();

    const tbody = document.querySelector('#leaderboard tbody');
    tbody.innerHTML = '';

    data.forEach((item, index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${index + 1}</td>
            <td>${item.user_id}</td>
            <td>${item.score}</td>
        `;
        tbody.appendChild(row);
    });
}

// Tải leaderboard mỗi 5 giây
loadLeaderboard();
setInterval(loadLeaderboard, 5000);