#!/user/bin/env ruby
require 'rubygems'
require 'rack'
require_relative 'decorator'
require_relative 'builder'

app = Builder.new {
  use Rack::ContentLength
  use Decorator, :header => "****************header****************</br>"
  run lambda {|env| [200, {"Content-Type"=>"text/html"}, ['hello world']]}
}.to_app

Rack::Handler::WEBrick.run app, :Port=>3000
