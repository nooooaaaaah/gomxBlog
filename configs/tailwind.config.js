module.exports = {
  content: [
    "./ui/**/*.{html,js}", // Correct glob pattern to include HTML and JS files
    "./static/js/**/*.js",
  ],
  theme: {
    extend: {
      // Custom colors
      colors: {
        midnight: {
          800: "#0d231e",
          900: "#0c1f1a",
          950: "#081411",
        },
        purp: {
          400: "#8c98f1",
        },
        syan: {
          200: "#1ed0bf",
          400: "#5aacc4",
        },
        peach: {
          200: "#febe83",
          600: "#ff6347",
        },
        "light-text": "#e8f9f7",
        "border-color": "#2c2f33",
        keyword: "#8c98f1",
        string: "#e5c76d",
        comment: "#616876",
        "function-keyword": "#5aacc4",
        tag: "#99ffe4",
      },
      borderWidth: {
        1: "1px",
      },
      fontStyle: {
        italic: "italic",
      },
      translate: {
        "extra-full": "150%",
        "sidebar-visible": "calc(100% - 24px)",
      },
      // Custom animation for sidebar opening and closing
      keyframes: {
        gopher: {
          "0%": { transform: "translateY(-102%)" },
          "100%": { transform: "translateY(-35%)" },
        },
        animateBackground: {
          "0%": { backgroundPosition: "0% 50%" },
          "50%": { backgroundPosition: "100% 50%" },
          "100%": { backgroundPosition: "0% 50%" },
        },
        slideIn: {
          "0%": { transform: "translateX(-100%)" },
          "100%": { transform: "translateX(0)" },
        },
        slideOut: {
          "0%": { transform: "translateX(0)" },
          "100%": { transform: "translateX(calc(-100% + 24px))" }, // Leave button visible
        },
        "chyron-scroll": {
          "0%": { transform: "translateX(100vw)" },
          "100%": { transform: "translateX(-100vw)" },
        },
        tentShiftRight: {
          "0%": { transform: "translateX(0)" },
          "100%": { transform: "translateX(40px)" }, // Shift right by the sidebar width
        },
        contentShiftLeft: {
          "0%": { transform: "translateX(40px)" }, // Assuming your sidebar is reversed here
          "100%": { transform: "translateX(0)" },
        },
        text: {
          "0%, 100%": {
            backgroundSize: "200% 200%",
            backgroundPosition: "left center",
          },
          "50%": {
            backgroundSize: "200% 200%",
            backgroundPosition: "right center",
          },
        },
        fadeIn: {
          "0%": { opacity: 0 },
          "100%": { opacity: 1 },
        },
        colorShift: {
          "0%, 100%": { backgroundColor: "#ff6347", color: "#ffdab9" },
          "25%": { backgroundColor: "#40e0d0", color: "#ffefd5" },
          "50%": { backgroundColor: "#ff8c00", color: "#4682b4" },
          "75%": { backgroundColor: "#da70d6", color: "#6495ed" },
        },
        bolden: {
          "0%": { transform: "scale(.4) translateX(-105%)" },
          "100%": { transform: "scale(1.3) translateX(10%)" },
        },
        maxWidth: {
          "80ch": "80ch",
          "60ch": "60ch",
          "40ch": "40ch",
        },
      },
      animation: {
        "chyron-scroll-ssl": "chyron-scroll 45s linear infinite",
        "chyron-scroll-sl": "chyron-scroll 35s linear infinite",
        "chyron-scroll-md": "chyron-scroll 25s linear infinite",
        "chyron-scroll-fs": "chyron-scroll 15s linear infinite",
        gopher: "gopher 5s ease-out forwards",
        slideIn: "slideIn 0.5s ease-out forwards",
        slideOut: "slideOut 0.5s ease-in forwards",
        contentShiftRight: "contentShiftRight 0.5s ease-out forwards",
        contentShiftLeft: "contentShiftLeft 0.5s ease-in forwards",
        text: "text 10s ease infinite bolden 3s infinite alternate ease-in-out",
        "fade-in": "fadeIn 2s ease-out",
        colorShift: "colorShift 8s infinite alternate ease-in-out",
        bolden: "bolden 3s ease-in forwards",
      },
    },
  },
  plugins: [],
};
