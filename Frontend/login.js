document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("login-form");
  
    form.addEventListener("submit", async (e) => {
      e.preventDefault();
  
      const data = {
        email: form.email.value,
        password: form.password.value,
      };
  
      try {
        const res = await fetch("http://localhost:8081/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        });
  
        const result = await res.json();
  
        if (res.ok) {
          alert(result.message || "Login berhasil!");
          // Misal redirect ke halaman utama
          // window.location.href = "/home.html";
        } else {
          alert(result.message || "Email atau password salah.");
        }
      } catch (err) {
        console.error("Error:", err);
        alert("Gagal terhubung ke server.");
      }
    });
  });
  