module.exports = {
	publicPath: './',
	/* 修改html标题 */
	chainWebpack: config => {
		config.plugin('html')
			.tap(args => {
				console.log(args);
				args[0].title = "中华诗词";
				return args;
			})
	},
}