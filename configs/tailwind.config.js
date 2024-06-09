// tailwind.config.js
module.exports = {
  content: ["./ui/**/*.{html,js}", "./static/js/**/*.js"],
  theme: {
    extend: {
      colors: {
        midnight: {
          50: "#f5faf9", // lightest shade
          100: "#e0f1ef", // lighter shade
          200: "#b3ded6", // light shade
          300: "#80c8bb", // moderate light shade
          400: "#4da19d", // moderate shade
          500: "#267c78", // default shade (mid-tone)
          600: "#1b615f", // dark shade
          700: "#124946", // darker shade
          800: "#0d231e", // darkest shade before almost black
          900: "#0c1f1a", // almost black
          950: "#081411", // near black
        },
        purp: {
          50: "#f6f7ff", // lightest shade
          100: "#eceeff", // lighter shade
          200: "#d6dbff", // light shade
          300: "#b6bbff", // moderate light shade
          400: "#8c98f1", // default shade (mid-tone)
          500: "#6a7ae6", // moderate dark shade
          600: "#5b67d4", // dark shade
          700: "#4a53b7", // darker shade
          800: "#394091", // darkest shade before almost black
          900: "#2b306f", // almost black
          950: "#1f224b", // near black
        },
        syan: {
          50: "#e5f9f7", // lightest shade
          100: "#ccf5f2", // lighter shade
          200: "#a9f0ea", // light shade
          300: "#78d7d0", // moderate light shade
          400: "#5aacc4", // default shade (mid-tone)
          500: "#47a5ae", // moderate dark shade
          600: "#2e857e", // dark shade
          700: "#236663", // darker shade
          800: "#1d514c", // darkest shade before almost black
          900: "#17423c", // almost black
          950: "#102e28", // near black
        },
        peach: {
          50: "#fff4e9", // lightest shade
          100: "#ffe7cc", // lighter shade
          200: "#febe83", // light shade
          300: "#ffab70", // moderate light shade
          400: "#ff9e61", // moderate shade
          500: "#ff8554", // default shade (mid-tone)
          600: "#ff6347", // moderate dark shade
          700: "#cc5139", // dark shade
          800: "#99402c", // darkest shade before almost black
          900: "#66301f", // almost black
          950: "#4d2316", // near black
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
      translate: {
        "extra-full": "150%",
        "sidebar-visible": "calc(100% - 24px)",
      },
      keyframes: {
        grow: {
          "0%": {
            transform: "scale(0.75)",
            opacity: "0",
          },
          "100%": {
            transform: "scale(1)",
            opacity: "1",
          },
        },
        slideUp: {
          "0%": { transform: "translateY(20px)", opacity: 0 },
          "100%": { transform: "translateY(0)", opacity: 1 },
        },
        gopher: {
          "0%": { transform: "translateY(-102%)", opacity: 0 },
          "5%": { transform: "translateY(-97%)", opacity: 0.1 },
          "10%": { transform: "translateY(-90%)", opacity: 0.2 },
          "15%": { transform: "translateY(-85%)", opacity: 0.25 },
          "20%": { transform: "translateY(-80%)", opacity: 0.3 },
          "25%": { transform: "translateY(-75%)", opacity: 0.35 },
          "30%": { transform: "translateY(-70%)", opacity: 0.4 },
          "35%": { transform: "translateY(-65%)", opacity: 0.45 },
          "40%": { transform: "translateY(-60%)", opacity: 0.5 },
          "45%": { transform: "translateY(-55%)", opacity: 0.55 },
          "50%": { transform: "translateY(-50%)", opacity: 0.6 },
          "55%": { transform: "translateY(-47%)", opacity: 0.65 },
          "60%": { transform: "translateY(-45%)", opacity: 0.7 },
          "65%": { transform: "translateY(-43%)", opacity: 0.75 },
          "70%": { transform: "translateY(-40%)", opacity: 0.8 },
          "75%": { transform: "translateY(-38%)", opacity: 0.85 },
          "80%": { transform: "translateY(-37%)", opacity: 0.9 },
          "85%": { transform: "translateY(-36%)", opacity: 0.95 },
          "90%": { transform: "translateY(-35%)", opacity: 0.97 },
          "95%": { transform: "translateY(-35%)", opacity: 0.98 },
          "100%": { transform: "translateY(-35%)", opacity: 1 },
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
          "100%": { transform: "translateX(calc(-100% + 24px))" },
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
          "0%": { transform: "scale(0.4) translateX(-105%)" },
          "100%": { transform: "scale(1.3) translateX(10%)" },
        },
      },
      animation: {
        grow: "grow 1s ease-out forwards",
        gopher: "gopher 5s ease-in 60s forwards", // Corrected animation with delay
        slideIn: "slideIn 0.5s ease-out forwards",
        slideOut: "slideOut 0.5s ease-in forwards",
        text: "text 10s ease infinite",
        "fade-in": "fadeIn 2s ease-out",
        colorShift: "colorShift 8s infinite alternate ease-in-out",
        bolden: "bolden 3s ease-in forwards",
        slideUp: "slideUp 2s ease-in-out forwards",
      },
      maxWidth: {
        "80ch": "80ch",
        "60ch": "60ch",
        "40ch": "40ch",
      },
      boxShadow: {
        sm: "0 1px 2px 0 rgba(0, 0, 0, 0.05)",
        md: "0 4px 6px -1px rgba(0, 0, 0, 0.1)",
      },
      backdropFilter: {
        none: "none",
        blur: "blur(10px)",
      },
      backgroundOpacity: {
        75: "0.75",
      },
    },
    plugins: [
      require("tailwindcss-filters"), // Ensure to install tailwindcss-filters plugin
    ],
  },
};
