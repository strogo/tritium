
require 'yaml'
require 'minitest/unit'
require_relative '../../lib/tritium/engines/debug/engine'
require_relative 'engine_tests'

class DebugEngineTest < MiniTest::Unit::TestCase
  include EngineTests
  include Tritium::Engines
  
  def engine_class
    Debug
  end
  
  def test_debug
    log = Logger.new(STDOUT)
    log.level = Logger::ERROR
    engine = engine_class.new("doc('xml') { $('.//body') {\n  select('.//a') { remove() } \n}\n}", :path => "/scripts", :logger => log)
    result, export_vars = engine.run("<body><div><a></a></div></body>")
    doc = Nokogiri::XML.parse(result)
    
    debug_output = engine.root_step.debug

    # Make sure we removed the a tag!
    assert_equal 2, doc.search("*").size

    #assert_equal "select", step.instruction.name
    #assert_equal "select", step.children.first.instruction.name
  end
end