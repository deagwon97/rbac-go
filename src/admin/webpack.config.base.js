module.exports = {
  stats: { modules: false },
  resolve: {
    extensions: [".js", ".jsx"],
    alias: {
      src: path.resolve(__dirname, `src`),
    },
  },
};
