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
        "extra-full": "150%", // Example of adding a custom translate value
        "sidebar-visible": "calc(100% - 24px)", // Adjust for button visibility
      },
      // Custom animation for sidebar opening and closing
      keyframes: {
        slideIn: {
          "0%": { transform: "translateX(-100%)" },
          "100%": { transform: "translateX(0)" },
        },
        slideOut: {
          "0%": { transform: "translateX(0)" },
          "100%": { transform: "translateX(calc(-100% + 24px))" }, // Leave button visible
        },
        contentShiftRight: {
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
        spin: {
          "0%": { transform: "rotate(0deg)" },
          "100%": { transform: "rotate(360deg)" },
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
          "0%": { transform: "scale(1)" },
          "100%": { transform: "scale(1.3)" },
        },
        maxWidth: {
          "80ch": "80ch",
          "60ch": "60ch",
          "40ch": "40ch",
        },
      },
      animation: {
        spin: "spin 1s linear infinite",
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
