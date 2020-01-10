from fcbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Fcbeat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        fcbeat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("fcbeat is running"))
        exit_code = fcbeat_proc.kill_and_wait()
        assert exit_code == 0
