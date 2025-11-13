class Tblang < Formula
  desc "Plugin-based Infrastructure as Code language"
  homepage "https://github.com/SwanHtetAungPhyo/tblang"
  url "https://github.com/SwanHtetAungPhyo/tblang/archive/refs/tags/v0.1.1.tar.gz"
  sha256 "079450df92dbb1b0f5b24658476789ad6ec4a1fdd27b5222b29343a1d046f67b"
  license "MIT"
  version "0.1.1"
  head "https://github.com/SwanHtetAungPhyo/tblang.git", branch: "main"

  depends_on "go" => :build

  def install
    if (buildpath/"core").exist?
      cd "core" do
        system "go", "build", "-ldflags", "-s -w", "-o", "tblang", "./cmd/tblang"
        bin.install "tblang"
      end

      cd "plugin/aws" do
        system "go", "build", "-o", "tblang-provider-aws", "main.go"
        (lib/"tblang/plugins").install "tblang-provider-aws"
      end

      if (buildpath/"tblang-demo").exist?
        (prefix/"examples").install Dir["tblang-demo/*.tbl"]
        (prefix/"examples").install Dir["tblang-demo/*.md"]
      end
    else
      odie "Repository structure not found. Please check the installation."
    end
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
