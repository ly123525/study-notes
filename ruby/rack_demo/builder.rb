class Builder
  def initialize(&block)
    @middlewares = []
    self.instance_eval(&block)
  end

  def use(middleware_class, *options, &block)
    @middlewares << lambda {|app| middleware_class.new(app, *options, &block)}
  end

  def run(app)
    @app = app
  end

  def to_app
    @middlewares.reverse.inject(@app) {|app, middleware| middleware.call(app)}
  end
end
