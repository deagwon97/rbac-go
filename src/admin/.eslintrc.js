module.exports = {
  root: true,
  env: {
    browser: true,
    es6: true,
  },
  plugins: ["react"],
  extends: [
    // typescript 표준 규칙 모음
    "eslint:recommended",
    "plugin:import/errors",
    "plugin:import/warnings",
    // "plugin:react/recommended",
    // import 관련 규칙 모음

    "plugin:prettier/recommended",
    // "prettier",
    // prettier 관련 규칙 모음
  ],
  settings: {
    "import/resolver": {
      node: {
        extensions: [".js", ".jsx", ".ts", ".tsx"],
        moduleDirectory: ["node_modules", "src/"],
      },
      webpack: {
        config: "webpack.config.base.js",
      },
    },
  },
  rules: {
    // 추가하고 싶은 rule을 더 추가해줍니다.
    "no-unused-vars": "warn",
    // "no-console": "off",
  },
};
