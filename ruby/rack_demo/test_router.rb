#!/user/bin/env ruby
require 'rubygems'
require 'rack'
# app = Rack::Builder.new {

#   map '/hello' do
#     run lambda { |env| [200, {'Content-Type' => 'text/html'}, ['hello']] }
#   end
#   map '/world' do
#     run lambda { |env| [200, {'Content-Type' => 'text/html'}, ['world']] }
#   end
#   map '/' do
#     run lambda { |env| [200, {'Content-Type' => 'text/html'}, ['all']] }
#   end
# }.to_app

app = Rack::Builder.new {
  use Rack::ContentLength
  map '/hello' do
    run lambda { |env| [200, {'Content-Type' => 'text/html'}, ["SCRIPT_NAME=#{env['SCRIPT_NAME']}", "PATH_INFO=#{env['PATH_INFO']}"]] }
  end
  map '/world' do
    run lambda { |env| [200, {'Content-Type' => 'text/html'}, ['world']] }
  end
}.to_app

Rack::Handler::WEBrick.run app, :Port => 3000

