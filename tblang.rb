class Tblang < Formula
  desc "Plugin-based Infrastructure as Code language"
  homepage "https://github.com/yourusername/tblang"
  url "https://github.com/yourusername/tblang/archive/refs/tags/v0.1.0.tar.gz"
  sha256 "REPLACE_WITH_ACTUAL_SHA256"
  license "MIT"
  version "0.1.0"

  depends_on "go" => :build

  def install
    cd "core" do
      system "go", "build", "-ldflags", "-s -w", "-o", "tblang", "./cmd/tblang"
      bin.install "tblang"
    end

    cd "plugin/aws" do
      system "go", "build", "-o", "tblang-provider-aws", "main.go"
      (lib/"tblang/plugins").install "tblang-provider-aws"
    end

    (prefix/"examples").install Dir["tblang-demo/*.tbl"]
    (prefix/"examples").install Dir["tblang-demo/*.md"]
  end

  def post_install
    (var/"tblang").mkpath
  end

  def caveats
    <<~EOS
      TBLang has been installed!

      The AWS provider plugin is installed at:
        #{lib}/tblang/plugins/tblang-provider-aws

      Example files are available at:
        #{prefix}/examples/

      To get started:
        1. Configure your AWS credentials:
           aws configure --profile your-profile

        2. Create a .tbl file (see examples)

        3. Run TBLang commands:
           tblang plan infrastructure.tbl
           tblang apply infrastructure.tbl
           tblang show
           tblang destroy infrastructure.tbl

      For more information:
        tblang --help
    EOS
  end

  test do
    system "#{bin}/tblang", "--version"
    system "#{bin}/tblang", "plugins", "list"
  end
end
