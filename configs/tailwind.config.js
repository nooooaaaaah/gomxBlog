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
          "100%": { transform: "translateX(calc(20rem - 24px))" }, // Shift right by the sidebar width
        },
        contentShiftLeft: {
          "0%": { transform: "translateX(calc(20rem - 24px))" }, // Assuming your sidebar is reversed here
          "100%": { transform: "translateX(0)" },
        },
      },
      animation: {
        slideIn: "slideIn 0.5s ease-out forwards",
        slideOut: "slideOut 0.5s ease-in forwards",
        contentShiftRight: "contentShiftRight 0.5s ease-out forwards",
        contentShiftLeft: "contentShiftLeft 0.5s ease-in forwards",
      },
    },
  },
  plugins: [],
};
