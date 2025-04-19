document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("register-form");

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const data = {
      first_name: form.first_name.value,
      last_name: form.last_name.value,
      email: form.email.value,
      password: form.password.value,
      confirm_password: form.confirm_password.value,
      phone: form.phone.value,
    };

    console.log(data);

    // Validasi password dan konfirmasi password
    if (form.password.value !== form.confirm_password.value) {
      alert("Password dan konfirmasi tidak sama!");
      return;
    }

    try {
      const res = await fetch("http://localhost:8081/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      // Cek apakah response statusnya sukses (2xx)
      if (res.ok) {
        // Jika sukses, baca response JSON
        const result = await res.json();
        alert(result.message || "Registrasi berhasil!");
        window.location.href = "./login.html"; // Redirect ke halaman login
      } else {
        // Jika response status tidak sukses, coba untuk baca respons sebagai teks
        const errorText = await res.text();
        try {
          // Coba parsing errorText sebagai JSON
          const errorResult = JSON.parse(errorText);
          alert(errorResult.message || "Terjadi kesalahan saat registrasi.");
        } catch (jsonError) {
          // Jika gagal parsing sebagai JSON, tampilkan sebagai teks biasa
          alert(errorText || "Terjadi kesalahan saat registrasi.");
        }
      }
    } catch (err) {
      console.error("Error:", err);
      alert("Gagal terhubung ke server. Pastikan server berjalan dengan baik.");
    }
  });
});
